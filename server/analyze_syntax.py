import pylint
import subprocess

class AnalyzeStatic():
    def __init__(self, file_path):
        self.file_path = file_path
        self.file_type = self.get_file_type()
        # Supported languages and their respective linters
        if self.file_type == "py":
            self.linter = self.run_pylint
        elif self.file_type == "js" or self.file_type == "jsx" or self.file_type == "ts":
            self.linter = self.run_eslint
        elif self.file_type == "cpp":
            self.linter = self.run_cppcheck
        else:
            self.linter = None

    def get_file_type(self):
        return self.file_path.split('.')[-1]

    def run_pylint(self):
        pylint_opts = [self.file_path]
        pylint_output = pylint.lint.Run(pylint_opts, do_exit=False)
        return pylint_output.linter.stats

    def run_eslint(self):
        result = subprocess.run(['eslint', self.file_path, '-f', 'json'], capture_output=True, text=True)
        return result.stdout

    def run_cppcheck(self):
        result = subprocess.run(['cppcheck', '--enable=all', '--xml', '--xml-version=2', self.file_path], capture_output=True, text=True)
        return result.stdout

    def lint(self):
        if self.linter:
            return self.linter()
        else:
            return "Unsupported file type"


