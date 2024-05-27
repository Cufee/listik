-- CreateTable
CREATE TABLE "users" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "profilePicture" TEXT,
    "locked" BOOLEAN NOT NULL DEFAULT false
);

-- CreateTable
CREATE TABLE "sessions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "cookieValue" TEXT NOT NULL,
    "identifier" TEXT NOT NULL,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "lastUsed" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "label" TEXT,
    "userId" TEXT NOT NULL,
    CONSTRAINT "sessions_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "groups" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "desc" TEXT
);

-- CreateTable
CREATE TABLE "group_members" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "groupId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "permissions" TEXT NOT NULL,
    CONSTRAINT "group_members_groupId_fkey" FOREIGN KEY ("groupId") REFERENCES "groups" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "group_members_userId_fkey" FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "group_invites" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "expiresAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdById" TEXT NOT NULL,
    "useLimit" INTEGER NOT NULL DEFAULT 1,
    "useCount" INTEGER NOT NULL DEFAULT 0,
    "enabled" BOOLEAN NOT NULL DEFAULT true,
    "code" TEXT NOT NULL,
    CONSTRAINT "group_invites_createdById_fkey" FOREIGN KEY ("createdById") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "item_tags" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "desc" TEXT,
    "icon" TEXT,
    "color" TEXT NOT NULL DEFAULT '#ffffff',
    "createdById" TEXT NOT NULL,
    "groupId" TEXT NOT NULL,
    CONSTRAINT "item_tags_createdById_fkey" FOREIGN KEY ("createdById") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "item_tags_groupId_fkey" FOREIGN KEY ("groupId") REFERENCES "groups" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "lists" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "desc" TEXT,
    "groupId" TEXT NOT NULL,
    "createdById" TEXT NOT NULL,
    CONSTRAINT "lists_groupId_fkey" FOREIGN KEY ("groupId") REFERENCES "groups" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "lists_createdById_fkey" FOREIGN KEY ("createdById") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "list_items" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "createdById" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "desc" TEXT,
    "price" TEXT,
    "quantity" INTEGER,
    "checked" BOOLEAN NOT NULL DEFAULT false,
    "listId" TEXT NOT NULL,
    CONSTRAINT "list_items_createdById_fkey" FOREIGN KEY ("createdById") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "list_items_listId_fkey" FOREIGN KEY ("listId") REFERENCES "lists" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "_redeemedBy" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL,
    CONSTRAINT "_redeemedBy_A_fkey" FOREIGN KEY ("A") REFERENCES "group_invites" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "_redeemedBy_B_fkey" FOREIGN KEY ("B") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "_ItemTagToListItem" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL,
    CONSTRAINT "_ItemTagToListItem_A_fkey" FOREIGN KEY ("A") REFERENCES "item_tags" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "_ItemTagToListItem_B_fkey" FOREIGN KEY ("B") REFERENCES "list_items" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateIndex
CREATE INDEX "users_locked_idx" ON "users"("locked");

-- CreateIndex
CREATE UNIQUE INDEX "sessions_cookieValue_key" ON "sessions"("cookieValue");

-- CreateIndex
CREATE INDEX "sessions_userId_idx" ON "sessions"("userId");

-- CreateIndex
CREATE INDEX "sessions_cookieValue_idx" ON "sessions"("cookieValue");

-- CreateIndex
CREATE INDEX "sessions_lastUsed_idx" ON "sessions"("lastUsed" DESC);

-- CreateIndex
CREATE INDEX "group_members_userId_idx" ON "group_members"("userId");

-- CreateIndex
CREATE INDEX "group_members_groupId_idx" ON "group_members"("groupId");

-- CreateIndex
CREATE UNIQUE INDEX "group_invites_code_key" ON "group_invites"("code");

-- CreateIndex
CREATE INDEX "group_invites_createdById_idx" ON "group_invites"("createdById");

-- CreateIndex
CREATE INDEX "group_invites_code_enabled_idx" ON "group_invites"("code", "enabled");

-- CreateIndex
CREATE INDEX "item_tags_groupId_idx" ON "item_tags"("groupId");

-- CreateIndex
CREATE INDEX "item_tags_createdById_idx" ON "item_tags"("createdById");

-- CreateIndex
CREATE INDEX "lists_groupId_idx" ON "lists"("groupId");

-- CreateIndex
CREATE INDEX "lists_createdById_idx" ON "lists"("createdById");

-- CreateIndex
CREATE INDEX "lists_updatedAt_idx" ON "lists"("updatedAt" DESC);

-- CreateIndex
CREATE INDEX "list_items_checked_idx" ON "list_items"("checked");

-- CreateIndex
CREATE INDEX "list_items_createdById_idx" ON "list_items"("createdById");

-- CreateIndex
CREATE INDEX "list_items_listId_idx" ON "list_items"("listId");

-- CreateIndex
CREATE UNIQUE INDEX "_redeemedBy_AB_unique" ON "_redeemedBy"("A", "B");

-- CreateIndex
CREATE INDEX "_redeemedBy_B_index" ON "_redeemedBy"("B");

-- CreateIndex
CREATE UNIQUE INDEX "_ItemTagToListItem_AB_unique" ON "_ItemTagToListItem"("A", "B");

-- CreateIndex
CREATE INDEX "_ItemTagToListItem_B_index" ON "_ItemTagToListItem"("B");
