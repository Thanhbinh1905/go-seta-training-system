-- Users
CREATE TABLE users (
  user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  role TEXT CHECK (role IN ('manager', 'member')) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- Teams
CREATE TABLE teams (
    team_id SERIAL PRIMARY KEY,
    team_name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE team_members (
    team_id INTEGER REFERENCES teams(team_id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE team_managers (
    team_id INTEGER REFERENCES teams(team_id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (team_id, user_id)
);

-- Folders
CREATE TABLE folders (
  folder_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  owner_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- Notes
CREATE TABLE notes (
  note_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title TEXT NOT NULL,
  body TEXT,
  folder_id UUID REFERENCES folders(folder_id) ON DELETE CASCADE,
  owner_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- Folder shares
CREATE TABLE folder_shares (
  folder_id UUID REFERENCES folders(folder_id) ON DELETE CASCADE,
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  access TEXT CHECK (access IN ('read', 'write')) NOT NULL,
  PRIMARY KEY (folder_id, user_id)
);

-- Note shares
CREATE TABLE note_shares (
  note_id UUID REFERENCES notes(note_id) ON DELETE CASCADE,
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  access TEXT CHECK (access IN ('read', 'write')) NOT NULL,
  PRIMARY KEY (note_id, user_id)
);
