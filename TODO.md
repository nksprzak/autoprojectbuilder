
## TODO:
   - Add CLI interface with github.com/spf13/cobra module
      - Add a 'version' command
      - Add a 'create' command for new projects
      - Add a 'modify' command for existing projects (Subtasks listed below)
   - ~~Add support for multiple authors on a single project~~
   - Add support for modifying existing projects
      - Adding/Removing contributors
      - Changing repo location (Remote origin URL)
      - Changing the project name
      - Changing the local directory of the project
   - ~~Add support for initializing cargo with Rust projects~~
   - Add support for getting author email addresses from the user
   - Add support for interactively getting README and TODO file contents from the user
   - Add support for getting git usernames from the user
   - Add support for getting the service that will host the git repository (Github, Gitlab, Bitbucket, etc.)
      - Add support for setting the upstream URL after running `git init` in the new project
   - Add support for C++
   - ~~Add better error handling support using the github.com/pkg/errors module~~
   - ~~Print the version of the tool with the name~~
