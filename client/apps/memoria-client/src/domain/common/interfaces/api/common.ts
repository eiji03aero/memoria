export type Paginate = {
  page: number;
  perPage: number;
};

export type Pagination = {
  per_page: number;
  current_page: number;
  total_page: number;
};

export type CPaginate = {
  cursor?: string;
  cbefore?: number;
  cafter?: number;
  cexclude?: boolean;
};

export type CPagination = {
  next_cursor: string;
  prev_cursor: string;
};
