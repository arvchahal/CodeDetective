// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';
import * as child_process from 'child_process';
// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "codedetective" is now active!');

	// The command has been defined in the package.json file
	// Now provide the implementation of the command with registerCommand
	// The commandId parameter must match the command field in package.json
	const disposable = vscode.commands.registerCommand('codedetective.helloWorld', () => {
		// The code you place here will be executed every time your command is executed
		// Display a message box to the user
		const editor = vscode.window.activeTextEditor;
		if (!editor) {
			return;
		}
		const document = editor.document;
		const filePath = document.fileName;
		const analyze_comm = `python3 server/analyze_loc.py ${filePath}`;
		child_process.exec(analyze_comm, (err, stdout, stderr) => {
			if (err) {
				console.error(err);
				return;
			}
			//need to parse the output and show where the errors are in the file
			const lines = stdout.split('\n');
			
		});
		const analyze_syntax = child_process.spawn('python3', ['server/analyze_syntax.py', filePath]);



		vscode.window.showInformationMessage('Hello World from CodeDetective!');
	});

	context.subscriptions.push(disposable);
}

// This method is called when your extension is deactivated
export function deactivate() {
	
}
