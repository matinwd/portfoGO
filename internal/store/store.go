package store

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	db *sql.DB
}

type Settings map[string]string

type ShowcaseItem struct {
	Title          string
	Meta           string
	BodyHTML       string
	WhyHTML        string
	DiagramMermaid sql.NullString
}

type SkillGroup struct {
	Title       string
	ContentHTML string
	AnchorID    sql.NullString
}

type Badge struct {
	Image string
	Alt   string
}

type BlogPost struct {
	Title          string
	PublishedLabel string
	Summary        string
}

type ResearchItem struct {
	Title   string
	Slug    string
	Meta    string
	Summary string
}

type ResearchPage struct {
	Title       string
	Slug        string
	Meta        string
	Description string
	ContentHTML string
}

type Experience struct {
	Company     string
	Role        string
	Description string
	Years       string
}

func Open(databaseURL string) (*Store, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) EnsureSchema(ctx context.Context) error {
	return s.execSQLFile(ctx, filepath.Join("db", "schema.sql"))
}

func (s *Store) SeedIfEmpty(ctx context.Context) error {
	var count int
	if err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM site_settings").Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return s.execSQLFile(ctx, filepath.Join("db", "seed.sql"))
}

func (s *Store) execSQLFile(ctx context.Context, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, string(data))
	return err
}

func (s *Store) GetSettings(ctx context.Context) (Settings, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT key, value FROM site_settings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	settings := Settings{}
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		settings[key] = value
	}
	return settings, rows.Err()
}

func (s *Store) ListAboutParagraphs(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT content FROM about_paragraphs ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, err
		}
		items = append(items, content)
	}
	return items, rows.Err()
}

func (s *Store) ListShowcaseItems(ctx context.Context) ([]ShowcaseItem, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT title, meta, body, why, diagram_mermaid FROM showcase_items ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShowcaseItem
	for rows.Next() {
		var item ShowcaseItem
		if err := rows.Scan(&item.Title, &item.Meta, &item.BodyHTML, &item.WhyHTML, &item.DiagramMermaid); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *Store) ListSkillGroups(ctx context.Context) ([]SkillGroup, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT title, content, anchor_id FROM skills_groups ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SkillGroup
	for rows.Next() {
		var item SkillGroup
		if err := rows.Scan(&item.Title, &item.ContentHTML, &item.AnchorID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *Store) ListTrustBadges(ctx context.Context) ([]Badge, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT image, alt FROM trust_badges ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Badge
	for rows.Next() {
		var item Badge
		if err := rows.Scan(&item.Image, &item.Alt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *Store) ListBlogPosts(ctx context.Context) ([]BlogPost, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT title, published_label, summary FROM blog_posts ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BlogPost
	for rows.Next() {
		var item BlogPost
		if err := rows.Scan(&item.Title, &item.PublishedLabel, &item.Summary); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *Store) ListResearchItems(ctx context.Context) ([]ResearchItem, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT title, slug, meta, summary FROM research_items ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResearchItem
	for rows.Next() {
		var item ResearchItem
		if err := rows.Scan(&item.Title, &item.Slug, &item.Meta, &item.Summary); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (s *Store) GetResearchPage(ctx context.Context, slug string) (ResearchPage, error) {
	var page ResearchPage
	row := s.db.QueryRowContext(ctx, "SELECT title, slug, meta, description, content_html FROM research_pages WHERE slug = $1", slug)
	if err := row.Scan(&page.Title, &page.Slug, &page.Meta, &page.Description, &page.ContentHTML); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ResearchPage{}, sql.ErrNoRows
		}
		return ResearchPage{}, err
	}
	return page, nil
}

func (s *Store) ListExperiences(ctx context.Context) ([]Experience, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT company, role, description, years FROM experiences ORDER BY position")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Experience
	for rows.Next() {
		var item Experience
		if err := rows.Scan(&item.Company, &item.Role, &item.Description, &item.Years); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
