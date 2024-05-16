import * as Joi from 'joi';

export const envSchema = Joi.object({
  PORT: Joi.string().required(),
  DATABASE_URL: Joi.string().required(),
});

export type EnvAcceptedValues = 'PORT' | 'DATABASE_URL';

export const envMsgs = {
  missingKey: 'Provided env file missing following params',
};
