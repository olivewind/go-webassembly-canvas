const go = new Go();

WebAssembly.instantiateStreaming(fetch("wasm/lib.wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);
    });