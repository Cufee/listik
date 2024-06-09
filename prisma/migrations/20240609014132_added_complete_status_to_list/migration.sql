-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_lists" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "desc" TEXT,
    "complete" BOOLEAN NOT NULL DEFAULT false,
    "groupId" TEXT NOT NULL,
    "createdById" TEXT NOT NULL,
    CONSTRAINT "lists_groupId_fkey" FOREIGN KEY ("groupId") REFERENCES "groups" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "lists_createdById_fkey" FOREIGN KEY ("createdById") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_lists" ("createdAt", "createdById", "desc", "groupId", "id", "name", "updatedAt") SELECT "createdAt", "createdById", "desc", "groupId", "id", "name", "updatedAt" FROM "lists";
DROP TABLE "lists";
ALTER TABLE "new_lists" RENAME TO "lists";
CREATE INDEX "lists_groupId_idx" ON "lists"("groupId");
CREATE INDEX "lists_createdById_idx" ON "lists"("createdById");
CREATE INDEX "lists_updatedAt_idx" ON "lists"("updatedAt" DESC);
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
