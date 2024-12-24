  <div align="center">
  <h1>Templ8</h1>

  [![GitHub last commit](https://img.shields.io/github/last-commit/1shubham7/templ8)](#)
  ![GitHub language count](https://img.shields.io/github/languages/count/1shubham7/templ8)
  ![GitHub top language](https://img.shields.io/github/languages/top/1shubham7/templ8)
  [![codecov](https://codecov.io/gh/1Shubham7/templ8/graph/badge.svg?token=mGg6p0S7KL)](https://codecov.io/gh/1Shubham7/templ8)

</div>

```
▀▀█▀▀ ▒█▀▀▀ ▒█▀▄▀█ ▒█▀▀█ ▒█░░░ ▒█▀▀▀█
░▒█░░ ▒█▀▀▀ ▒█▒█▒█ ▒█▄▄█ ▒█░░░ █░▄▄▄█░
░▒█░░ ▒█▄▄▄ ▒█░░▒█ ▒█░░░ ▒█▄▄█ █▄▄▄▄█░

🚀 Transform any GitHub repository into your project template

templ8 simplifies project bootstrapping by creating templates from existing GitHub repos.
Just provide a repository URL and branch name - templ8 handles the rest.

Usage:
  templ8 <github-repo-url> -b=<branch-name> -d=<distination>

Happy coding! ✨
```

Templ8 is a CLI tool that helps you bootstrap projects quickly by transforming any GitHub repository into a template. With just the repository URL and branch name, Templ8 takes care of the rest, saving you time on repetitive setup tasks.

Whether you're starting a new project or creating a template for your team, Templ8 makes it easy to clone and customize GitHub repositories with a few commands.

⭐ Star us on GitHub — it motivates me a lot!

[![Share](https://img.shields.io/badge/share-000000?logo=x&logoColor=white)](https://x.com/intent/tweet?text=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/1Shubham7/templ8%20%23OpenIDConnect%20%23Security%20%23Authentication)
[![Share](https://img.shields.io/badge/share-1877F2?logo=facebook&logoColor=white)](https://www.facebook.com/sharer/sharer.php?u=https://github.com/1Shubham7/templ8)
[![Share](https://img.shields.io/badge/share-0A66C2?logo=linkedin&logoColor=white)](https://www.linkedin.com/sharing/share-offsite/?url=https://github.com/1Shubham7/templ8)
[![Share](https://img.shields.io/badge/share-FF4500?logo=reddit&logoColor=white)](https://www.reddit.com/submit?title=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/1Shubham7/templ8)
[![Share](https://img.shields.io/badge/share-0088CC?logo=telegram&logoColor=white)](https://t.me/share/url?url=https://github.com/1Shubham7/templ8&text=Check%20out%20this%20project%20on%20GitHub)

## 📚Tech Stack

- **Language:** Golang
- **Tools:** Cobra CLI
- **APIs:** GitHub API
- **CI Pipeline:** GitHub Actions

## Installing

Follow these steps to install and set up Templ8:

1. Visit the [Releases](https://github.com/1Shubham7/templ8/releases) section of the Templ8 repository and download the latest stable release.

2. Open your shell configuration file (e.g., `~/.bashrc` or `~/.zshrc`) in a text editor:
   ```bash
   nano ~/.bashrc
   ```

3. Add the following line at the end of the file to include the directory in your PATH (replace `PATH_TO_EXECUTABLE` with the actual path to your downloaded executable):
   ```bash
   export PATH=$PATH:[PATH_TO_EXECUTABLE]
   ```

4. Save and close the file:
   - For nano: Press `Ctrl+O`, then `Enter`, and `Ctrl+X`.

5. Reload the shell configuration:
   ```bash
   source ~/.bashrc
   ```

6. Navigate to the directory where you placed the CLI executable:
   ```bash
   cd [PATH_TO_EXECUTABLE]
   ```

7. Make the executable file runnable:
   ```bash
   chmod +x templ8.exe
   ```

8. Verify the installation by running the following command:
   ```bash
   templ8.exe --help
   ```

9. Additionally you can also rename the templ8.exe to templ8:
    ```bash
    mv templ8.exe templ8
    ```

You're all set to transform any GitHub repository into a template! 🎉

## Usage

Simply call the executable with the URL of the repository you want to use as a template, and the tool will download the latest version of the specified branch into the directory you want and initialize it as a new repository.

```bash
templ8 https://github.com/1Shubham7/demo-repo-for-templ8
```

### Optional Arguments

You can specify additional options such as a different branch or output directory:

- **Specify a branch:** Use the `--branch` or `-b` flag to specify a branch to clone.
- **Set an output folder:** Use the `-distination` or `-d` flag to set the name of the folder where the contents will be placed.

Example:

```bash
# This will create a folder called 'new-proj' with the contents of the source repo, specifically the 'dev' branch.
templ8 -b=dev -d=new-proj https://github.com/1Shubham7/demo-repo-for-templ8
```
