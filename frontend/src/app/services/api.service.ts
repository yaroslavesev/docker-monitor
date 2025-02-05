/**
 * ApiService for handling HTTP calls
 */
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Container } from '../models/container';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  private apiUrl = environment.apiBaseUrl + '/api';

  constructor(private http: HttpClient) {}

  //getContainers - gets all containers
  getContainers(): Observable<Container[]> {
    return this.http.get<Container[]>(`${this.apiUrl}/containers`);
  }

  //createContainer -  creates a new container
  createContainer(data: { ip_address: string }): Observable<Container> {
    return this.http.post<Container>(`${this.apiUrl}/containers`, data);
  }

  //updateContainer -  updates container by ID
  updateContainer(id: number, payload: Partial<Container>): Observable<Container> {
    return this.http.put<Container>(`${this.apiUrl}/containers/${id}`, payload);
  }

  //deleteContainer - deletes container by ID
  deleteContainer(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/containers/${id}`);
  }
}
