import { Injectable } from '@nestjs/common';
import { CreateSongPayload } from './songs.validations';
import { PrismaService } from 'src/common';
import { Prisma } from '@prisma/client';

@Injectable()
export class SongsRepository {
  constructor(private prisma: PrismaService) {}
  async createSong(createSongPayload: CreateSongPayload) {
    // get existing user ids
    const existUserIdResult = await this.prisma.artists.findMany({
      where: {
        userId: {
          in: createSongPayload.artists,
        },
      },
      select: {
        userId: true,
      },
    });

    // prepare existing user id array
    const existingArtistsIds: number[] = ((results: { userId: number }[]) => {
      const returnData = [] as number[];
      results.map((result) => {
        returnData.push(result.userId);
      });
      return returnData;
    })(existUserIdResult);

    // prepare createArtistsData to connectOrCreate
    const createArtistsData: Prisma.ArtistsSongsCreateWithoutSongInput[] = ((
      artistsIds: number[],
    ) => {
      const artistsData = [] as Prisma.ArtistsSongsCreateWithoutSongInput[];
      artistsIds.forEach((id: number) => {
        if (existingArtistsIds.includes(id)) {
          // if id is registered as artists, then it just connect it with the song
          artistsData.push({
            artist: {
              connect: {
                userId: id,
              },
            },
          });
        } else {
          //register artists id with song
          artistsData.push({
            artist: {
              create: {
                userId: id,
              },
            },
          });
        }
      });
      return artistsData;
    })(createSongPayload.artists);

    return this.prisma.songs.create({
      data: {
        duration: createSongPayload.releasedDate,
        releasedDate: createSongPayload.releasedDate,
        title: createSongPayload.title,
        lyrics: createSongPayload.lyrics,
        artists: {
          create: createArtistsData,
        },
      },
      include: {
        artists: true,
      },
    });
  }
}
