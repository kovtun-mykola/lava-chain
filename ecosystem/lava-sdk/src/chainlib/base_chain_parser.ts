import {
  functionTag,
  ParseDirective,
  ApiCollection,
  Api,
  Header,
  Header_HeaderType,
} from "../codec/lavanet/lava/spec/api_collection";
import { Metadata } from "../codec/lavanet/lava/pairing/relay";
import { Spec } from "../codec/lavanet/lava/spec/spec";
import { Logger } from "../logger/logger";

// const APIInterfaceJsonRPC = "jsonrpc";
// const APIInterfaceTendermintRPC = "tendermintrpc";
const APIInterfaceRest = "rest";
// const APIInterfaceGrpc = "grpc";

interface ApiKey {
  name: string;
  connectionType: string;
}

type ApiKeyString = string;

function ApiKeyToString(key: ApiKey): ApiKeyString {
  return JSON.stringify(key);
}

interface TaggedContainer {
  parsing: ParseDirective;
  apiCollection: ApiCollection;
}

interface CollectionKey {
  connectionType: string;
  internalPath: string;
  addon: string;
}

interface ApiContainer {
  api: Api;
  collectionKey: CollectionKey;
  apiKey: ApiKey;
}

interface HeaderContainer {
  header: Header;
  apiKey: ApiKey;
}

interface VerificationKey {
  extension: string;
  addon: string;
}

interface VerificationContainer {
  connectionType: string;
  name: string;
  parseDirective: ParseDirective;
  value: string;
  latestDistance: Long;
  verificationKey: VerificationKey;
}

type VerificationKeyString = string;

function VerificationKeyToString(key: VerificationKey): VerificationKeyString {
  return JSON.stringify(key);
}

export class BaseChainParser {
  private taggedApis: Map<functionTag, TaggedContainer>;
  private spec: Spec;
  private serverApis: Map<ApiKeyString, ApiContainer>;
  private headers: Map<ApiKeyString, HeaderContainer>;
  private apiCollections: Map<CollectionKey, ApiCollection>;
  // TODO: implement addons.
  private allowedAddons: Set<string>;
  // private extensionParser: ExtensionParser;
  private apiInterface: string;
  private verifications: Map<VerificationKeyString, VerificationContainer[]>;

  constructor(spec: Spec, apiInterface: string) {
    this.apiInterface = apiInterface;
    this.taggedApis = new Map();
    this.serverApis = new Map();
    this.apiCollections = new Map();
    this.headers = new Map();
    this.allowedAddons = new Set();
    this.verifications = new Map();
    this.spec = spec;

    if (spec.enabled) {
      for (const apiCollection of spec.apiCollections) {
        if (!apiCollection.enabled) {
          continue;
        }
        if (apiCollection.collectionData?.apiInterface != this.apiInterface) {
          continue;
        }
        const collectionKey: CollectionKey = {
          connectionType: apiCollection.collectionData.type,
          internalPath: apiCollection.collectionData.internalPath,
          addon: apiCollection.collectionData.addOn,
        };

        // parse directives
        for (const parsing of apiCollection.parseDirectives) {
          this.taggedApis.set(parsing.functionTag, {
            parsing: parsing,
            apiCollection: apiCollection,
          });
        }

        // parse api collection
        for (const api of apiCollection.apis) {
          if (!api.enabled) {
            continue;
          }
          let apiName = api.name;
          if (this.apiInterface == APIInterfaceRest) {
            const re = /{[^}]+}/;
            apiName = api.name.replace(re, "replace-me-with-regex");
            apiName = apiName.replace(/replace-me-with-regex/g, "[^\\/\\s]+");
            apiName = this.escapeRegExp(apiName); // Assuming you have a RegExp.escape function
          }
          const apiKey: ApiKey = {
            name: apiName,
            connectionType: collectionKey.connectionType,
          };
          this.serverApis.set(ApiKeyToString(apiKey), {
            apiKey: apiKey,
            api: api,
            collectionKey: collectionKey,
          });
        }

        // Parse headers
        for (const header of apiCollection.headers) {
          const apiKeyHeader: ApiKey = {
            name: header.name,
            connectionType: collectionKey.connectionType,
          };
          this.headers.set(ApiKeyToString(apiKeyHeader), {
            header: header,
            apiKey: apiKeyHeader,
          });
        }

        for (const verification of apiCollection.verifications) {
          for (const parseValue of verification.values) {
            const verificationKey: VerificationKey = {
              extension: parseValue.extension,
              addon: apiCollection.collectionData.addOn,
            };
            if (!verification.parseDirective) {
              throw new Error(
                Logger.fatal(
                  "Missing verification parseDirective data in BaseChainParser constructor",
                  verification
                )
              );
            }
            const verificationContainer: VerificationContainer = {
              connectionType: apiCollection.collectionData.type,
              name: verification.name,
              parseDirective: verification.parseDirective,
              value: parseValue.expectedValue,
              latestDistance: parseValue.latestDistance,
              verificationKey: verificationKey,
            };
            const vfkey = VerificationKeyToString(verificationKey);
            const existingVerifications = this.verifications.get(vfkey);
            if (!existingVerifications) {
              this.verifications.set(vfkey, [verificationContainer]);
            } else {
              existingVerifications.push(verificationContainer);
            }
          }
        }
      }
    }
  }

  handleHeaders(
    metadata: Metadata[],
    apiCollection: ApiCollection,
    headersDirection: Header_HeaderType
  ): HeadersHandler {
    if (metadata.length == 0) {
      return {
        filteredHeaders: [],
        overwriteRequestedBlock: "",
        ignoredMetadata: [],
      };
    }
    const retMetaData: Metadata[] = [];
    const ignoredMetadata: Metadata[] = [];
    let overwriteRequestedBlock = "";
    for (const header of metadata) {
      const headerName = header.name.toLowerCase();
      if (!apiCollection.collectionData) {
        throw new Error(
          Logger.fatal(
            "Missing api collection data in handleHeaders",
            apiCollection
          )
        );
      }
      const apiKey: ApiKey = {
        name: headerName,
        connectionType: apiCollection.collectionData?.type,
      };

      const headerDirective = this.headers.get(ApiKeyToString(apiKey));
      if (!headerDirective) {
        continue; // this header is not handled
      }
      if (
        headerDirective.header.kind == headersDirection ||
        headerDirective.header.kind == Header_HeaderType.pass_both
      ) {
        retMetaData.push(header);
        if (
          headerDirective.header.functionTag ==
          functionTag.SET_LATEST_IN_METADATA
        ) {
          overwriteRequestedBlock = header.value;
        }
      } else if (headerDirective.header.kind == Header_HeaderType.pass_ignore) {
        ignoredMetadata.push(header);
      }
    }

    return {
      filteredHeaders: retMetaData,
      ignoredMetadata: ignoredMetadata,
      overwriteRequestedBlock: overwriteRequestedBlock,
    };
  }

  isAddon(addon: string): boolean {
    return this.allowedAddons.has(addon);
  }

  escapeRegExp(s: string): string {
    return s.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  }

  matchSpecApiByName(
    name: string,
    connectionType: string
  ): [ApiContainer | undefined, boolean] {
    let foundNameOnDifferentConnectionType: string | undefined = undefined;
    for (const [, api] of this.serverApis.entries()) {
      const re = new RegExp(`^${api.apiKey.name}$`);
      if (re.test(name)) {
        if (api.apiKey.connectionType === connectionType) {
          return [api, true];
        } else {
          foundNameOnDifferentConnectionType = api.apiKey.connectionType;
        }
      }
    }
    if (foundNameOnDifferentConnectionType) {
      Logger.warn(
        `Found the api on a different connection type, found: ${foundNameOnDifferentConnectionType}, requested: ${connectionType}`
      );
    }
    return [undefined, false];
  }
}

interface HeadersHandler {
  filteredHeaders: Metadata[];
  overwriteRequestedBlock: string;
  ignoredMetadata: Metadata[];
}
