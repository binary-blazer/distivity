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

3. Rename the `.env.example` file to `.env` and update the values:

   ```sh
   # On Unix
   cp .env.example .env

   # On Windows
   ren .env.example .env
   ```
   
   The `.env` file should look like this with your values:
   ```env
   PORT=8080 # Port
   DISCORD_BOT_TOKEN="" # Bot token
   DISCORD_GUILD_ID="" # Guild ID
   DISCORD_GUILD_INVITE="" # Guild invite
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