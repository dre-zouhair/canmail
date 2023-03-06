db = db.getSiblingDB("mailer");

db.createCollection("targets");
db.createCollection("templates");

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