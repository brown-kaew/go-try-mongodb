db.createUser({
    user: "admin",
    pwd: "pass",
    roles: [
        {
            role: "readWrite",
            db: "mydb",
        },
    ],
});

db = new Mongo().getDB("mydb");

db.createCollection("user", { capped: false });

db.user.insert([
    { id: 1, name: "Kaew" },
    { id: 2, name: "Arin" },
    { id: 3, name: "John" },
]);
