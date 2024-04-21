-- CreateTable
CREATE TABLE "Users" (
    "id" SERIAL NOT NULL,
    "firstName" TEXT NOT NULL,
    "lastName" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,

    CONSTRAINT "Users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Songs" (
    "id" SERIAL NOT NULL,
    "title" TEXT NOT NULL,
    "releasedDate" TIMESTAMP(3) NOT NULL,
    "duration" TIMESTAMP(3) NOT NULL,
    "lyrics" TEXT,

    CONSTRAINT "Songs_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Playlists" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "ownerId" INTEGER NOT NULL,

    CONSTRAINT "Playlists_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Artists" (
    "userId" INTEGER NOT NULL,

    CONSTRAINT "Artists_pkey" PRIMARY KEY ("userId")
);

-- CreateTable
CREATE TABLE "ArtistsSongs" (
    "songId" INTEGER NOT NULL,
    "artistId" INTEGER NOT NULL,

    CONSTRAINT "ArtistsSongs_pkey" PRIMARY KEY ("songId","artistId")
);

-- AddForeignKey
ALTER TABLE "Playlists" ADD CONSTRAINT "Playlists_ownerId_fkey" FOREIGN KEY ("ownerId") REFERENCES "Users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Artists" ADD CONSTRAINT "Artists_userId_fkey" FOREIGN KEY ("userId") REFERENCES "Users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ArtistsSongs" ADD CONSTRAINT "ArtistsSongs_songId_fkey" FOREIGN KEY ("songId") REFERENCES "Songs"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ArtistsSongs" ADD CONSTRAINT "ArtistsSongs_artistId_fkey" FOREIGN KEY ("artistId") REFERENCES "Artists"("userId") ON DELETE RESTRICT ON UPDATE CASCADE;
