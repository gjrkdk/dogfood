import { listNotes } from "@/lib/db";

// Server component: leest bij elke request uit Postgres. Bewust geen caching —
// dit is een demo-app om te bewijzen dat de agent tegen een draaiende database
// kan werken, niet een productie-frontend.
export const dynamic = "force-dynamic";

export default async function Home() {
  const notes = await listNotes();

  return (
    <main style={{ fontFamily: "monospace", padding: "2rem" }}>
      <h1>Notes</h1>
      {notes.length === 0 ? (
        <p>Nog geen notities.</p>
      ) : (
        <ul>
          {notes.map((note) => (
            <li key={note.id}>{note.body}</li>
          ))}
        </ul>
      )}
    </main>
  );
}
