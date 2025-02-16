# Distivity

Distivity is a modern RESTful API that allows you to get presence details of a Discord account by ID. It provides endpoints to fetch user information, avatars, banners, and status.

## API Capabilities

### Endpoints

- **GET /**: Returns basic information about the API.
- **GET /user/:id**: Fetches detailed information about a user by their Discord ID.
- **GET /avatar/:id**: Retrieves the avatar of a user by their Discord ID.
- **GET /banner/:id**: Retrieves the banner of a user by their Discord ID.
- **GET /status**: Returns the status of the API.

### Features

- Fetch user information including username, discriminator, avatar, banner, and activity.
- Retrieve user avatars and banners in high resolution.
- Get the current status of the API.

## Self-Hosting Requirements and Setup Instructions

### Requirements

- Go 1.23.3 or higher
- A Discord bot token
- A server or local machine to run the API

### Setup Instructions

1. Clone the repository:

   ```sh
   git clone https://github.com/binary-blazer/distivity.git
   cd distivity
   ```

2. Update the configuration in `./config/variables.go`

3. Create a `.env` file in the root directory and add your Discord bot token:

   ```sh
   DISCORD_BOT_TOKEN=your_bot_token_here
   ```

4. Install the dependencies:

   ```sh
   go mod tidy
   ```

5. Run the API:

   ```sh
   go run main.go
   ```

6. The API should now be running on `http://localhost:3000`.

## Code of Conduct

### Our Pledge

In the interest of fostering an open and welcoming environment, we as contributors and maintainers pledge to make participation in our project and our community a harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Our Standards

Examples of behavior that contributes to creating a positive environment include:

- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

- The use of sexualized language or imagery and unwelcome sexual attention or advances
- Trolling, insulting/derogatory comments, and personal or political attacks
- Public or private harassment
- Publishing others' private information, such as a physical or electronic address, without explicit permission
- Other conduct which could reasonably be considered inappropriate in a professional setting

### Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable behavior and are expected to take appropriate and fair corrective action in response to any instances of unacceptable behavior.

### Scope

This Code of Conduct applies both within project spaces and in public spaces when an individual is representing the project or its community. Examples of representing a project or community include using an official project e-mail address, posting via an official social media account, or acting as an appointed representative at an online or offline event. Representation of a project may be further defined and clarified by project maintainers.

### Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported by contacting the project team at [INSERT EMAIL ADDRESS]. All complaints will be reviewed and investigated and will result in a response that is deemed necessary and appropriate to the circumstances. The project team is obligated to maintain confidentiality with regard to the reporter of an incident. Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good faith may face temporary or permanent repercussions as determined by other members of the project's leadership.

## Security Policy

### Reporting a Vulnerability

If you discover a security vulnerability within this project, please send an email to [INSERT EMAIL ADDRESS]. All security vulnerabilities will be promptly addressed.

## Collaboration and Contribution Instructions

We welcome contributions from the community! To get started, follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch from the `main` branch.
3. Make your changes and commit them with clear and concise messages.
4. Push your changes to your forked repository.
5. Open a pull request to the `main` branch of the original repository.

Please ensure that your code adheres to the project's coding standards and passes all tests before submitting a pull request.

Thank you for your contributions!
