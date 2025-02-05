/**
 * Container model interface
 */
export interface Container {
  id: number;
  ip_address: string;
  last_ping_time: string;
  last_success_time: string;
  created_at: string;
  updated_at: string;
}
