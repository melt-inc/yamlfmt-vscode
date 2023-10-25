const assert = require('assert');

// You can import and use all API from the 'vscode' module
// as well as import your extension to test it
const vscode = require('vscode');
// const myExtension = require('../extension');

const wait = ms => new Promise(resolve => setTimeout(resolve, ms));

suite('Extension Test Suite', () => {
	vscode.window.showInformationMessage('Start all tests.');

	test('Default compact sequence style', async () => {
		// close all open editors
		await vscode.commands.executeCommand('workbench.action.closeAllEditors');

		assert.equal(vscode.workspace.getConfiguration('yaml').get('compactSequenceStyle'), true);
	});

	test('Format document', async () => {
		// close all open editors
		await vscode.commands.executeCommand('workbench.action.closeAllEditors');

		// setup doc
		doc = await vscode.workspace.openTextDocument({ content: 'hello:\n  - you\n  - world\n  - universe\n' });
		await vscode.languages.setTextDocumentLanguage(doc, 'yaml');

		// show it
		await vscode.window.showTextDocument(doc, vscode.ViewColumn.Active, false);
		// set tab size to 2
		await vscode.workspace.getConfiguration('editor').update('tabSize', 2, vscode.ConfigurationTarget.Global);
		// format document
		await vscode.commands.executeCommand('editor.action.formatDocument');

		// get doc content as a string
		const docContent = doc.getText();
		console.log(docContent);

		assert.equal(docContent, 'hello:\n- you\n- world\n- universe');
	});
});
