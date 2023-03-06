
db.use("mailer")

db.createUser({
    user: "mailer",
    pwd: "password",
    roles: [
        {
            role: "readWrite",
            db: "mailer"
        }
    ]
});

db.createCollection("targets");
db.createCollection("templates");