package sqlite

// const (
// 	project = `CREATE TABLE PROJECT(
// 		ID INTEGER PRIMARY KEY AUTOINCREMENT,
// 		NAME STRING NOT NULL,
// 		MANAGERID INT NOT NULL

// 	)`
// 	question = `CREATE TABLE QUESTION(
// 		ID INTEGER PRIMARY KEY AUTOINCREMENT,
// 		[ORDER] INTEGER NOT NULL,
// 		QUESTION STRING NOT NULL,
// 		ANSWER STRING NOT NULL,
// 		PROJECTID INT NOT NULL
// 	)    `
// 	user = `CREATE TABLE USER (
// 		ID        INTEGER     PRIMARY KEY
// 							  AUTOINCREMENT,
// 		USERNAME  STRING (25) UNIQUE
// 							  NOT NULL,
// 		PROJECTID INTEGER     REFERENCES PROJECT (ID),
//		CHATID 	  INTEGER     NOT NULL,
// 		ONCHAT    BOOLEAN     DEFAULT (FALSE),
// 		FIRSTNAME STRING (25),
// 		LASTNAME  STRING
// 	);`
// 	manager = `CREATE TABLE MANAGER (
// 		ID        INTEGER     PRIMARY KEY
// 							  AUTOINCREMENT,
// 		USERNAME  STRING (25) UNIQUE
// 							  NOT NULL,
// 		ISBUSY    BOOLEAN     DEFAULT (FALSE),
//		CHATID 	  INTEGER     NOT NULL,
// 		CURRENTCLIENTID INTEGER     REFERENCES USER (ID),
// 		FIRSTNAME STRING (25),
// 		LASTNAME  STRING
// 	);`
// )
