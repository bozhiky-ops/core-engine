/**
 * Core Engine Types
 * 
 * This file contains type definitions for the Core Engine project.
 */

export type Primitive = string | number | boolean | null | undefined;

export interface ApiResponse<T> {
  data: T;
  status: number;
  message?: string;
  timestamp: string;
}

export type Nullable<T> = T | null;

export type Optional<T> = T | undefined;

export type Dictionary<T> = {
  [key: string]: T;
};

export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';

export interface PaginatedResponse<T> {
  items: T[];
  total: number;
  page: number;
  limit: number;
  hasNext: boolean;
}

export type EventHandler<T = unknown> = (event: T) => void;

export type AsyncFunction<T = unknown, R = unknown> = (arg: T) => Promise<R>;

export type Constructor<T = unknown> = new (...args: any[]) => T;

export type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P];
};

export type Mutable<T> = {
  -readonly [P in keyof T]: T[P];
};

export type ValueOf<T> = T[keyof T];