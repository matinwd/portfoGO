CREATE TABLE IF NOT EXISTS site_settings (
  key TEXT PRIMARY KEY,
  value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS about_paragraphs (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS showcase_items (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  meta TEXT NOT NULL,
  body TEXT NOT NULL,
  why TEXT NOT NULL,
  diagram_mermaid TEXT,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS skills_groups (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  anchor_id TEXT,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS trust_badges (
  id SERIAL PRIMARY KEY,
  image TEXT NOT NULL,
  alt TEXT NOT NULL,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS blog_posts (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  published_label TEXT NOT NULL,
  summary TEXT NOT NULL,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS research_items (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  slug TEXT UNIQUE NOT NULL,
  meta TEXT NOT NULL,
  summary TEXT NOT NULL,
  position INT NOT NULL
);

CREATE TABLE IF NOT EXISTS research_pages (
  slug TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  meta TEXT NOT NULL,
  description TEXT NOT NULL,
  content_html TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS experiences (
  id SERIAL PRIMARY KEY,
  company TEXT NOT NULL,
  role TEXT NOT NULL,
  description TEXT NOT NULL,
  years TEXT NOT NULL,
  position INT NOT NULL
);
