CREATE TABLE user_group (
  id SERIAL PRIMARY KEY NOT NULL,
  name text NOT NULL
);

CREATE TABLE "user" (
  id SERIAL PRIMARY KEY NOT NULL,
  name text NOT NULL,
  mail text NOT NULL,
  group_id INT UNIQUE REFERENCES user_group(id)
  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE "order" (
  id SERIAL PRIMARY KEY NOT NULL,
  name text NOT NULL,
  user_id INT UNIQUE REFERENCES "user"(id)
  ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE product_detail (
  id SERIAL PRIMARY KEY NOT NULL,
  description text NOT NULL
);

CREATE TABLE product (
  id SERIAL PRIMARY KEY NOT NULL,
  name text NOT NULL,
  product_detail_id INT UNIQUE REFERENCES product_detail(id)
);

CREATE TABLE order_product(
  order_id INT UNIQUE REFERENCES "order" (id),
  product_id INT UNIQUE REFERENCES product(id),
  PRIMARY KEY (order_id, product_id)
);
