if (!globalThis.crypto) {
    const crypto = require("crypto");
    globalThis.crypto = {
        getRandomValues(b) {
            crypto.randomFillSync(b);
        },
    };
}
