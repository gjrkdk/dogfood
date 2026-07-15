import { Pool } from "pg";

// DATABASE_URL komt uit de omgeving: lokaal via .env, in een Polder-workspace
// via de env-vars van de repo (die schrijft polderd naar ~/workspace/.env en
// exporteert ze in de shell). Zonder database is er niets te doen — falen met
// een duidelijke boodschap is beter dan een lege lijst tonen.
function connectionString(): string {
  const url = process.env.DATABASE_URL;
  if (!url) {
    throw new Error("DATABASE_URL ontbreekt (zie docker-compose.yml)");
  }
  return url;
}

let pool: Pool | undefined;

export function db(): Pool {
  pool ??= new Pool({ connectionString: connectionString() });
  return pool;
}

export type Note = {
  id: number;
  body: string;
};

export async function addNote(body: string): Promise<Note> {
  const { rows } = await db().query<Note>(
    "INSERT INTO notes (body) VALUES ($1) RETURNING id, body",
    [body],
  );
  return rows[0];
}

export async function listNotes(): Promise<Note[]> {
  const { rows } = await db().query<Note>(
    "SELECT id, body FROM notes ORDER BY id DESC",
  );
  return rows;
}
