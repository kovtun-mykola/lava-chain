// source: lavanet/lava/plans/policy.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('../../../gogoproto/gogo_pb.js');
goog.object.extend(proto, gogoproto_gogo_pb);
var cosmos_base_v1beta1_coin_pb = require('../../../cosmos/base/v1beta1/coin_pb.js');
goog.object.extend(proto, cosmos_base_v1beta1_coin_pb);
var lavanet_lava_spec_api_collection_pb = require('../../../lavanet/lava/spec/api_collection_pb.js');
goog.object.extend(proto, lavanet_lava_spec_api_collection_pb);
goog.exportSymbol('proto.lavanet.lava.plans.ChainPolicy', null, global);
goog.exportSymbol('proto.lavanet.lava.plans.ChainRequirement', null, global);
goog.exportSymbol('proto.lavanet.lava.plans.Policy', null, global);
goog.exportSymbol('proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.plans.Policy = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.lavanet.lava.plans.Policy.repeatedFields_, null);
};
goog.inherits(proto.lavanet.lava.plans.Policy, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.lavanet.lava.plans.Policy.displayName = 'proto.lavanet.lava.plans.Policy';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.plans.ChainPolicy = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.lavanet.lava.plans.ChainPolicy.repeatedFields_, null);
};
goog.inherits(proto.lavanet.lava.plans.ChainPolicy, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.lavanet.lava.plans.ChainPolicy.displayName = 'proto.lavanet.lava.plans.ChainPolicy';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.plans.ChainRequirement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.lavanet.lava.plans.ChainRequirement.repeatedFields_, null);
};
goog.inherits(proto.lavanet.lava.plans.ChainRequirement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.lavanet.lava.plans.ChainRequirement.displayName = 'proto.lavanet.lava.plans.ChainRequirement';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.lavanet.lava.plans.Policy.repeatedFields_ = [1,7];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.plans.Policy.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.plans.Policy.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.plans.Policy} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.Policy.toObject = function(includeInstance, msg) {
  var f, obj = {
    chainPoliciesList: jspb.Message.toObjectList(msg.getChainPoliciesList(),
    proto.lavanet.lava.plans.ChainPolicy.toObject, includeInstance),
    geolocationProfile: jspb.Message.getFieldWithDefault(msg, 2, 0),
    totalCuLimit: jspb.Message.getFieldWithDefault(msg, 3, 0),
    epochCuLimit: jspb.Message.getFieldWithDefault(msg, 4, 0),
    maxProvidersToPair: jspb.Message.getFieldWithDefault(msg, 5, 0),
    selectedProvidersMode: jspb.Message.getFieldWithDefault(msg, 6, 0),
    selectedProvidersList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.plans.Policy}
 */
proto.lavanet.lava.plans.Policy.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.plans.Policy;
  return proto.lavanet.lava.plans.Policy.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.plans.Policy} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.plans.Policy}
 */
proto.lavanet.lava.plans.Policy.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.lavanet.lava.plans.ChainPolicy;
      reader.readMessage(value,proto.lavanet.lava.plans.ChainPolicy.deserializeBinaryFromReader);
      msg.addChainPolicies(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGeolocationProfile(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setTotalCuLimit(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setEpochCuLimit(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxProvidersToPair(value);
      break;
    case 6:
      var value = /** @type {!proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE} */ (reader.readEnum());
      msg.setSelectedProvidersMode(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.addSelectedProviders(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.plans.Policy.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.plans.Policy.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.plans.Policy} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.Policy.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getChainPoliciesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.lavanet.lava.plans.ChainPolicy.serializeBinaryToWriter
    );
  }
  f = message.getGeolocationProfile();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getTotalCuLimit();
  if (f !== 0) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = message.getEpochCuLimit();
  if (f !== 0) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = message.getMaxProvidersToPair();
  if (f !== 0) {
    writer.writeUint64(
      5,
      f
    );
  }
  f = message.getSelectedProvidersMode();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getSelectedProvidersList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      7,
      f
    );
  }
};


/**
 * repeated ChainPolicy chain_policies = 1;
 * @return {!Array<!proto.lavanet.lava.plans.ChainPolicy>}
 */
proto.lavanet.lava.plans.Policy.prototype.getChainPoliciesList = function() {
  return /** @type{!Array<!proto.lavanet.lava.plans.ChainPolicy>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.lavanet.lava.plans.ChainPolicy, 1));
};


/**
 * @param {!Array<!proto.lavanet.lava.plans.ChainPolicy>} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
*/
proto.lavanet.lava.plans.Policy.prototype.setChainPoliciesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.lavanet.lava.plans.ChainPolicy=} opt_value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.plans.ChainPolicy}
 */
proto.lavanet.lava.plans.Policy.prototype.addChainPolicies = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.lavanet.lava.plans.ChainPolicy, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.clearChainPoliciesList = function() {
  return this.setChainPoliciesList([]);
};


/**
 * optional int32 geolocation_profile = 2;
 * @return {number}
 */
proto.lavanet.lava.plans.Policy.prototype.getGeolocationProfile = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setGeolocationProfile = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional uint64 total_cu_limit = 3;
 * @return {number}
 */
proto.lavanet.lava.plans.Policy.prototype.getTotalCuLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setTotalCuLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional uint64 epoch_cu_limit = 4;
 * @return {number}
 */
proto.lavanet.lava.plans.Policy.prototype.getEpochCuLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setEpochCuLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional uint64 max_providers_to_pair = 5;
 * @return {number}
 */
proto.lavanet.lava.plans.Policy.prototype.getMaxProvidersToPair = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setMaxProvidersToPair = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional SELECTED_PROVIDERS_MODE selected_providers_mode = 6;
 * @return {!proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE}
 */
proto.lavanet.lava.plans.Policy.prototype.getSelectedProvidersMode = function() {
  return /** @type {!proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {!proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setSelectedProvidersMode = function(value) {
  return jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * repeated string selected_providers = 7;
 * @return {!Array<string>}
 */
proto.lavanet.lava.plans.Policy.prototype.getSelectedProvidersList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.setSelectedProvidersList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.addSelectedProviders = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.lavanet.lava.plans.Policy} returns this
 */
proto.lavanet.lava.plans.Policy.prototype.clearSelectedProvidersList = function() {
  return this.setSelectedProvidersList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.lavanet.lava.plans.ChainPolicy.repeatedFields_ = [2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.plans.ChainPolicy.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.plans.ChainPolicy} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.ChainPolicy.toObject = function(includeInstance, msg) {
  var f, obj = {
    chainId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    apisList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    requirementsList: jspb.Message.toObjectList(msg.getRequirementsList(),
    proto.lavanet.lava.plans.ChainRequirement.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.plans.ChainPolicy}
 */
proto.lavanet.lava.plans.ChainPolicy.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.plans.ChainPolicy;
  return proto.lavanet.lava.plans.ChainPolicy.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.plans.ChainPolicy} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.plans.ChainPolicy}
 */
proto.lavanet.lava.plans.ChainPolicy.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setChainId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addApis(value);
      break;
    case 3:
      var value = new proto.lavanet.lava.plans.ChainRequirement;
      reader.readMessage(value,proto.lavanet.lava.plans.ChainRequirement.deserializeBinaryFromReader);
      msg.addRequirements(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.plans.ChainPolicy.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.plans.ChainPolicy} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.ChainPolicy.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getChainId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getApisList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getRequirementsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.lavanet.lava.plans.ChainRequirement.serializeBinaryToWriter
    );
  }
};


/**
 * optional string chain_id = 1;
 * @return {string}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.getChainId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.setChainId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string apis = 2;
 * @return {!Array<string>}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.getApisList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.setApisList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.addApis = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.clearApisList = function() {
  return this.setApisList([]);
};


/**
 * repeated ChainRequirement requirements = 3;
 * @return {!Array<!proto.lavanet.lava.plans.ChainRequirement>}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.getRequirementsList = function() {
  return /** @type{!Array<!proto.lavanet.lava.plans.ChainRequirement>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.lavanet.lava.plans.ChainRequirement, 3));
};


/**
 * @param {!Array<!proto.lavanet.lava.plans.ChainRequirement>} value
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
*/
proto.lavanet.lava.plans.ChainPolicy.prototype.setRequirementsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.lavanet.lava.plans.ChainRequirement=} opt_value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.plans.ChainRequirement}
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.addRequirements = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.lavanet.lava.plans.ChainRequirement, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.lavanet.lava.plans.ChainPolicy} returns this
 */
proto.lavanet.lava.plans.ChainPolicy.prototype.clearRequirementsList = function() {
  return this.setRequirementsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.lavanet.lava.plans.ChainRequirement.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.plans.ChainRequirement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.plans.ChainRequirement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.ChainRequirement.toObject = function(includeInstance, msg) {
  var f, obj = {
    collection: (f = msg.getCollection()) && lavanet_lava_spec_api_collection_pb.CollectionData.toObject(includeInstance, f),
    extensionsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    mixed: jspb.Message.getBooleanFieldWithDefault(msg, 3, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.plans.ChainRequirement}
 */
proto.lavanet.lava.plans.ChainRequirement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.plans.ChainRequirement;
  return proto.lavanet.lava.plans.ChainRequirement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.plans.ChainRequirement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.plans.ChainRequirement}
 */
proto.lavanet.lava.plans.ChainRequirement.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new lavanet_lava_spec_api_collection_pb.CollectionData;
      reader.readMessage(value,lavanet_lava_spec_api_collection_pb.CollectionData.deserializeBinaryFromReader);
      msg.setCollection(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addExtensions(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setMixed(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.plans.ChainRequirement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.plans.ChainRequirement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.plans.ChainRequirement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCollection();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      lavanet_lava_spec_api_collection_pb.CollectionData.serializeBinaryToWriter
    );
  }
  f = message.getExtensionsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getMixed();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * optional lavanet.lava.spec.CollectionData collection = 1;
 * @return {?proto.lavanet.lava.spec.CollectionData}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.getCollection = function() {
  return /** @type{?proto.lavanet.lava.spec.CollectionData} */ (
    jspb.Message.getWrapperField(this, lavanet_lava_spec_api_collection_pb.CollectionData, 1));
};


/**
 * @param {?proto.lavanet.lava.spec.CollectionData|undefined} value
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
*/
proto.lavanet.lava.plans.ChainRequirement.prototype.setCollection = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.clearCollection = function() {
  return this.setCollection(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.hasCollection = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated string extensions = 2;
 * @return {!Array<string>}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.getExtensionsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.setExtensionsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.addExtensions = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.clearExtensionsList = function() {
  return this.setExtensionsList([]);
};


/**
 * optional bool mixed = 3;
 * @return {boolean}
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.getMixed = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.lavanet.lava.plans.ChainRequirement} returns this
 */
proto.lavanet.lava.plans.ChainRequirement.prototype.setMixed = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * @enum {number}
 */
proto.lavanet.lava.plans.SELECTED_PROVIDERS_MODE = {
  ALLOWED: 0,
  MIXED: 1,
  EXCLUSIVE: 2,
  DISABLED: 3
};

goog.object.extend(exports, proto.lavanet.lava.plans);
