package db

const (
  CREATE_TABLE_STREAMERS = `
    CREATE TABLE IF NOT EXISTS streamers (
      username VARCHAR(25) PRIMARY KEY,
      email VARCHAR(128) NOT NULL,
      display_name VARCHAR(25) NOT NULL,
      img_url TEXT NOT NULL,
      token UUID NOT NULL
    );
  `
  CREATE_TABLE_FOLLOWERS = `
    CREATE TABLE IF NOT EXISTS followers (
      email VARCHAR(128) PRIMARY KEY,
      display_name VARCHAR(25) NOT NULL UNIQUE,
      token UUID NOT NULL
    );
  `
  CREATE_TABLE_SUBSCRIPTIONS = `
    CREATE TABLE IF NOT EXISTS subscriptions (
      streamer_username VARCHAR(25) REFERENCES streamers(username),
      follower_email VARCHAR(128) REFERENCES followers(email),
      subscribed_at TIMESTAMP NOT NULL DEFAULT Now(),
      bending BOOLEAN DEFAULT true,
      PRIMARY KEY (streamer_username, follower_email)
    );
  `
  CREATE_TABLE_CHANNELS = `
    CREATE TABLE IF NOT EXISTS channels (
      id SERIAL PRIMARY KEY,
      streamer VARCHAR(25) REFERENCES streamers(username),
      name VARCHAR(45) NOT NULL
    );
  `
  CREATE_TABLE_SECTIONS = `
    CREATE TABLE IF NOT EXISTS sections (
      id SERIAL PRIMARY KEY,
      channel_id INT REFERENCES channels(id),
      name VARCHAR(45) NOT NULL
    );
  `
  CREATE_TABLE_VIDEOS = `
    CREATE TABLE IF NOT EXISTS videos (
      id SERIAL PRIMARY KEY,
      section_id INT REFERENCES sections(id),
      title VARCHAR(256) NOT NULL,
      url TEXT NOT NULL
    );
  `
  CREATE_TABLE_COMMENTS = `
    CREATE TABLE IF NOT EXISTS comments (
      video_id INT REFERENCES videos(id),
      follower_email VARCHAR(128) REFERENCES followers(email),
      content TEXT NOT NULL,
      commented_at TIMESTAMP NOT NULL DEFAULT Now()
    );
  `
  CREATE_TABLE_LIKES = `
    CREATE TABLE IF NOT EXISTS likes (
      video_id INT REFERENCES videos(id),
      follower_email VARCHAR(128) REFERENCES followers(email),
      PRIMARY KEY (video_id, follower_email)
    );
  `
)