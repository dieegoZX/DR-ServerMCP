# MCP Server API Documentation

This document provides instructions on how to connect to and interact with the MCP Server API.

## Base URL

The API is running on `http://localhost:8080`.

## Endpoints

### Health Check

- **GET /ping**
  - **Description:** Checks if the server is running.
  - **Response (200 OK):**
    ```
    pong
    ```

### MCP Context Management

#### Create a new MCP Context

- **POST /mcp**
  - **Description:** Creates a new MCP context.
  - **Request Body:**
    - **Content-Type:** `application/json`
    - **Body:** A JSON object of any structure.
      ```json
      {
        "name": "example",
        "value": 123
      }
      ```
  - **Response (201 Created):**
    - **Content-Type:** `application/json`
    - **Body:** The created MCP context object, including the server-generated `id` and `created_at` timestamp.
      ```json
      {
        "id": "...",
        "created_at": "...",
        "data": {
          "name": "example",
          "value": 123
        }
      }
      ```

#### Get all MCP Contexts

- **GET /mcp**
  - **Description:** Retrieves a list of all MCP contexts.
  - **Response (200 OK):**
    - **Content-Type:** `application/json`
    - **Body:** A JSON array of MCP context objects.
      ```json
      [
        {
          "id": "...",
          "created_at": "...",
          "data": { ... }
        }
      ]
      ```

#### Get a specific MCP Context

- **GET /mcp/{id}**
  - **Description:** Retrieves a single MCP context by its ID.
  - **URL Parameters:**
    - `id` (string, required): The ID of the context to retrieve.
  - **Response (200 OK):**
    - **Content-Type:** `application/json`
    - **Body:** The requested MCP context object.
  - **Response (404 Not Found):** If no context with the given ID is found.

#### Update an MCP Context

- **PUT /mcp/{id}**
  - **Description:** Updates the data of an existing MCP context.
  - **URL Parameters:**
    - `id` (string, required): The ID of the context to update.
  - **Request Body:**
    - **Content-Type:** `application/json`
    - **Body:** A JSON object with the new data for the context.
  - **Response (200 OK):**
    - **Content-Type:** `application/json`
    - **Body:** The updated MCP context object.
  - **Response (404 Not Found):** If no context with the given ID is found.

#### Delete an MCP Context

- **DELETE /mcp/{id}**
  - **Description:** Deletes an MCP context.
  - **URL Parameters:**
    - `id` (string, required): The ID of the context to delete.
  - **Response (204 No Content):** On successful deletion.
  - **Response (404 Not Found):** If no context with the given ID is found.
