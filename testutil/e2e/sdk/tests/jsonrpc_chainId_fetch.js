const { LavaSDK } = require("../../../../ecosystem/lava-sdk/bin/src/sdk/sdk");

async function main() {
    // Initialize Lava SDK
    const eth = await LavaSDK.create({
        privateKey: process.env.PRIVATE_KEY,
        chainIds: "ETH1",
        lavaChainId:"lava",
        pairingListConfig:process.env.PAIRING_LIST,
        allowInsecureTransport: true,
    }).catch(e => {
        throw new Error(" ERR failed initializing lava-sdk jsonrpc test");
    });

    let relayArray = [];
    for (let i = 0; i < 100; i++) { // send 100 relays asynchronously
        relayArray.push((async () => {
            // Fetch chain id
            const result = await eth.sendRelay({
                method: "eth_chainId",
                params: [],
            }).catch(e => {
                throw new Error(` ERR ${i} [jsonrpc_chainId_fetch] failed sending relay jsonrpc test`);
            });

            // Parse response
            const parsedResponse = result;

            const chainID = parsedResponse.result;

            // Validate chainID
            if (chainID != "0x1") {
                throw new Error(" ERR [jsonrpc_chainId_fetch] Chain ID is not equal to 0x1");
            } else{
                console.log(i, "[jsonrpc_chainId_fetch] Success: Fetching ETH chain ID using jsonrpc passed. Chain ID correctly matches '0x1'");
            }
        })().catch(err => {throw err;}));
    }
    // wait for all relays to finish;
    await Promise.allSettled(relayArray);
}

(async () => {
    try {
        await main();
        process.exit(0);
    } catch (error) {
        console.error(" ERR "+error.message);
        process.exit(1);
    }
})();