import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../services/api.service';
import { Container } from '../../models/container';
import { HttpClientModule } from '@angular/common/http';

@Component({
  selector: 'app-containers',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './containers.component.html',
  styleUrls: ['./containers.component.css']
})
export class ContainersComponent implements OnInit {
  containers: Container[] = [];
  newIp: string = '';

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.loadContainers();
  }

  loadContainers(): void {
    this.apiService.getContainers().subscribe({
      next: (data) => this.containers = data,
      error: (err) => console.error('Failed to get containers', err)
    });
  }

  createContainer(): void {
    if (!this.newIp) return;

    this.apiService.createContainer({ ip_address: this.newIp }).subscribe({
      next: (created) => {
        this.containers.push(created);
        this.newIp = '';  // Clear input
      },
      error: (err) => console.error('Failed to create container', err)
    });
  }
}
