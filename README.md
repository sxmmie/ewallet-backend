# System Requirements and Architecture

- Language: Go (Golang) for high-performance and concurrent transaction handling.
- Database: PostgreSQL or MySQL with ACID compliance to prevent balance discrepancies.

- Key Components
  - User Service: Manages user registration and authentication.
  - Wallet Service: Handles balance tracking and storage.
  - Transaction Service: Processes top-ups, withdrawals, and peer-to-peer transfers safely using database transactions.