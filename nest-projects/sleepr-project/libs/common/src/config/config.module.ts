import { Module } from '@nestjs/common';
import { ConfigModule as NestConfigModule } from '@nestjs/config';
import * as joi from 'joi';

export interface envVars {
  MONGODB_URI: string;
}

@Module({
  imports: [
    NestConfigModule.forRoot({
      validationSchema: joi.object({
        MONGODB_URI: joi.string().required(),
      }),
    }),
  ],
})
export class ConfigModule {}
