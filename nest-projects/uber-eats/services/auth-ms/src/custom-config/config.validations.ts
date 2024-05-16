import * as Joi from 'joi';

export const envSchema = Joi.object({
  PORT: Joi.string().required(),
});

export type EnvAcceptedValues = 'PORT';

export const envMsgs = {
  missingKey: 'Provided env file missing following params',
};
