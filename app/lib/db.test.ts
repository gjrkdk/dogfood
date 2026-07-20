import { afterAll, beforeEach, describe, expect, it } from "vitest";

import { addNote, db, deleteNote, listNotes, updateNote } from "./db";

// Deze test praat met een échte Postgres — geen mock. Dat is het punt: hij
// faalt als de database niet draait, en bewijst daarmee dat de omgeving
// (docker compose + DATABASE_URL) klopt vóór er code gemerged wordt.
describe("notes", () => {
  beforeEach(async () => {
    await db().query("TRUNCATE notes RESTART IDENTITY");
  });

  afterAll(async () => {
    await db().end();
  });

  it("bewaart een notitie en geeft hem terug", async () => {
    const created = await addNote("hallo polder");
    expect(created.id).toBeGreaterThan(0);
    expect(created.body).toBe("hallo polder");

    const notes = await listNotes();
    expect(notes).toHaveLength(1);
    expect(notes[0].body).toBe("hallo polder");
  });

  it("geeft de nieuwste notitie eerst", async () => {
    await addNote("eerste");
    await addNote("tweede");

    const notes = await listNotes();
    expect(notes.map((n) => n.body)).toEqual(["tweede", "eerste"]);
  });

  it("begint leeg", async () => {
    expect(await listNotes()).toEqual([]);
  });

  it("verwijdert een bestaande notitie", async () => {
    const created = await addNote("weg ermee");

    expect(await deleteNote(created.id)).toBe(true);
    expect(await listNotes()).toEqual([]);
  });

  it("geeft false voor een onbekend id", async () => {
    expect(await deleteNote(12345)).toBe(false);
  });

  it("wijzigt de tekst van een bestaande notitie", async () => {
    const created = await addNote("origineel");

    const updated = await updateNote(created.id, "aangepast");
    expect(updated).toEqual({ id: created.id, body: "aangepast" });

    const notes = await listNotes();
    expect(notes.map((n) => n.body)).toEqual(["aangepast"]);
  });

  it("geeft null voor een onbekend id bij updaten", async () => {
    expect(await updateNote(12345, "wat dan ook")).toBeNull();
  });
});
