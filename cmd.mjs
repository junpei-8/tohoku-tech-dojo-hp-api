#!/usr/bin/env npx -p zx -p dotenv -p dotenv-expand zx
'use strict';

const moduleDir = $.env._.split('/').slice(0, -2).join('/');
const env = require(`${moduleDir}/dotenv`).config({ path: `${__dirname}/.env` });
require(`${moduleDir}/dotenv-expand`).expand(env);

const CMD = {
  /** gcloud へログインする */
  'login-gcloud': () => CMD.loginGcloud(),
  async loginGcloud() {
    await $`gcloud auth login`;
  },

  /** 開発環境の DB にマイグレーションする */
  'migrate-dev': () => CMD.migrateDev(),
  async migrateDev() {
    const { DB_CONNECTION_URL } = $.env;
    await $([`atlas migrate apply --url ${DB_CONNECTION_URL} --dir file://migrations ${getAllArgs('url', 'dir')}`]);
  },

  /** api をデプロイする */
  'deploy-api': () => CMD.deployApi(),
  async deployApi() {
    const { GCP_PROJECT_ID, GCP_REGION, GCP_API_PROJECT_NAME, GCP_API_IMAGE, GCP_SQL_INSTANCE, GCP_SQL_URL } = $.env;
    const BUILD_CONFIG_FILE = '.cloudbuild.api.yml';
    const DOCKER_FILE = 'Dockerfile.api';
    const DOCKER_STAGE_NAME = 'runner';

    await $ol`
      gcloud builds submit
      --config ${BUILD_CONFIG_FILE}
      --substitutions=${oneLine`
        _DOCKER_FILE=${DOCKER_FILE},
        _DOCKER_STAGE_NAME=${DOCKER_STAGE_NAME},
        _PROJECT_ID=${GCP_PROJECT_ID},
        _PROJECT_NAME=${GCP_API_PROJECT_NAME},
        _REGION=${GCP_REGION},
        _IMAGE=${GCP_API_IMAGE},
        _SQL_INSTANCE=${GCP_SQL_INSTANCE},
        _SQL_URL=${GCP_SQL_URL}
      `}
      ${getAllArgs('config', 'substitutions')}
    `;
  },

  /** sql をデプロイする（まだ動作しない）*/
  'deploy-sql': () => CMD.deploySql(),
  async deploySql() {
    const { GCP_SQL_URL } = $.env;
    const BUILD_CONFIG_FILE = '.cloudbuild.sql.yml';

    await $ol`
      gcloud builds submit
      --config ${BUILD_CONFIG_FILE}
      --substitutions=_URL=${GCP_SQL_URL}
      ${getAllArgs('config', 'substitutions')}
    `;
  },
};

// 実行
const exec = CMD[argv._[0]];
if (!exec) throw '存在しないコマンドです。';

exec();

/**
 * 引数として受け取った値を１つの文字列として実行する
 * @props excludes {string[]} - 除外するキー
 */
function getAllArgs(...excludes) {
  excludes.push('_');

  return Object.keys(argv)
    .filter((key) => !excludes.includes(key))
    .reduce((acc, key) => {
      const arg = argv[key];
      if (arg === true) return ` ${acc} --${key}`;
      if (arg === false) return ` ${acc}`;
      if (Array.isArray(arg)) return ` ${acc} --${key} ${arg.join(' ')}`;
      return ` ${acc} --${key} ${arg}`;
    }, '');
}

/**
 * テンプレートリテラルを文字列へ変換する
 * @param {string[] | string} strings - テンプレートリテラルの文字列
 * @param {string[]} values - テンプレートリテラルに埋め込まれた変数の値
 */
function templateLiteralToString(strings, ...values) {
  // prettier-ignore
  return Array.isArray(strings)
    ? strings.reduce((acc, str, i) => `${acc}${str}${values[i] || ''}`, '')
    : strings;
}

/** 改行を削除し、空白を１つに連結させる */
function $ol(strings, ...values) {
  return $([
    templateLiteralToString(strings, ...values)
      .replace(/\n/g, ' ')
      .replace(/\s+/g, ' ')
      .trim(),
  ]);
}

/** 全ての改行と空白を削除する */
function oneLine(strings, ...values) {
  return templateLiteralToString(strings, ...values).replace(/\s/g, '');
}
