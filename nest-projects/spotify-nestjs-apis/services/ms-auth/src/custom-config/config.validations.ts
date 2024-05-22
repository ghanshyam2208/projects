import * as Joi from 'joi';

export const envSchema = Joi.object({
  HTTP_PORT: Joi.string().required(),
  DATABASE_URL: Joi.string().required(),
});

export type EnvAcceptedValues = 'HTTP_PORT' | 'DATABASE_URL';

export const envMsgs = {
  missingKey: 'Provided env file missing following params',
};
