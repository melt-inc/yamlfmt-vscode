// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
const vscode = require('vscode');
const read = require('fs').readFileSync;
const join = require('path').join;

require("./wasm_exec_polyfills");
require("./wasm_exec");
const go = new Go();

/**
 * this method is called when your extension is activated
 * @param {vscode.ExtensionContext} context
 */
async function activate(context) {

	const wasm_path = join(context.extensionPath, 'internal', 'yamlfmt.wasm');
	const wasm_bytes = read(wasm_path);

	// eslint-disable-next-line no-undef
	const waml = await WebAssembly.instantiate(wasm_bytes, go.importObject);
	go.run(waml.instance);

	let disposable = vscode.languages.registerDocumentRangeFormattingEditProvider('yaml', {
		provideDocumentRangeFormattingEdits: function (document, range, options) {
			// expand range to start/end of line
			range = range.with(document.lineAt(range.start.line).range.start, document.lineAt(range.end.line).range.end)

			// settings
			compact = vscode.workspace.getConfiguration('yaml').get('compactSequenceStyle')

			let oldText = document.getText(range);
			let newText = global.yamlfmt(oldText, options, compact);
			if (!newText) {
				newText = oldText
			}

			return [vscode.TextEdit.replace(range, newText)];
		}
	});

	context.subscriptions.push(disposable);
}

// this method is called when your extension is deactivated
function deactivate() { }

module.exports = {
	activate,
	deactivate
}
