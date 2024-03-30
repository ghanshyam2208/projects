export type MySqlConnectionType = {
  CONNECTION_STRING: string;
  DB: string;
  DBNAME: string;
};

export const MysqlDummyConnection: MySqlConnectionType = {
  CONNECTION_STRING: 'MYSQL://127.0.0.1/db', // sample conn
  DB: 'MYSQL',
  DBNAME: 'TEST',
};
