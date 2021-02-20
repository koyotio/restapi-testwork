CREATE TABLE IF NOT EXISTS "categories" (
    "id" SERIAL NOT NULL UNIQUE,
    "name" VARCHAR NOT NULL,
    "tags" JSONB NOT NULL,
    PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL NOT NULL UNIQUE,
    "category_id" INT NOT NULL,
    "articul" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT FK--products-category_id--categories-id
        FOREIGN KEY ("category_id")
            REFERENCES "categories"("id")
            ON DELETE CASCADE
            ON UPDATE CASCADE
);