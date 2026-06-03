CREATE TABLE USER (
        id VARCHAR(10) PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) NOT NULL,
        name VARCHAR(255),
        bio VARCHAR(500),
        avatarUrl VARCHAR(255),
        role VARCHAR(100) NOT NULL,
        githubUsername VARCHAR(20),
        linkedinUrl VARCHAR(25),
        portfolioUrl VARCHAR(25),
        city VARCHAR(15),
        totalPoints INT DEFAULT 0,
        seasonPoints INT DEFAULT 0,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        password VARCHAR(12) NOT NULL
);