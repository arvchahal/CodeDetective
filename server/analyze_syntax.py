import pylint.lint
import subprocess
import argparse
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
        print('Running pylint')
        pylint_opts = [self.file_path]
        try:
            pylint_output = pylint.lint.Run(pylint_opts)
            output = pylint_output.linter.stats
            self.results(output)
        except Exception as e:
            print(f"Error running pylint: {e}")

    def run_eslint(self):
        result = subprocess.run(['eslint', self.file_path, '-f', 'json'], capture_output=True, text=True)
        self.results(result)
        return result.stdout

    def run_cppcheck(self):
        result = subprocess.run(['cppcheck', '--enable=all', '--xml', '--xml-version=2', self.file_path], capture_output=True, text=True)
        self.results(result)
        return result

    def lint(self):
        if self.linter:
            return self.linter()
        else:
            return "Unsupported file type"
    def results(self,output):
        print(output)


def main():
    parser = argparse.ArgumentParser(description="Analyze syntax of a file")
    parser.add_argument("filepath", help="Path to file to be analyzed")
    args = parser.parse_args()
    analyzer = AnalyzeStatic(args.filepath)
    analyzer.lint()

if __name__ == '__main__':
    main()