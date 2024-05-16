import { Injectable, InternalServerErrorException } from '@nestjs/common';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';
import { envSchema, EnvAcceptedValues, envMsgs } from './config.validations';

@Injectable()
export class CustomConfigService {
  private envConfig: unknown;

  constructor() {}

  loadEnv() {
    // const options = { folder: './config' };

    const filePath = `${process.env.NODE_ENV || 'development'}.env`;
    const envFile = path.resolve(__dirname, '../../', filePath);

    // console.log('__dirname', __dirname);
    // console.log('options.folder', options.folder);
    // console.log('filePath', filePath);
    // console.log('envFile', envFile);
    // console.log(dotenv.parse(fs.readFileSync(envFile)));

    this.envConfig = dotenv.parse(fs.readFileSync(envFile));

    const { error } = envSchema.validate(this.envConfig, {
      abortEarly: false,
    });

    if (error) {
      throw new InternalServerErrorException(error, envMsgs.missingKey);
    }
    // console.log('this.envConfig', this.envConfig);
    return this.envConfig;
  }

  get(key: EnvAcceptedValues): string {
    return this.envConfig[key];
  }
}
