import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../services/api.service';
import { Container } from '../../models/container';

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
  private refreshIntervalId: any;

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.refreshIntervalId = setInterval(() => {
      this.loadContainers();
    }, 10_000);
    this.loadContainers();
  }

  ngOnDestroy(): void {
    if (this.refreshIntervalId) {
      clearInterval(this.refreshIntervalId);
    }
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

  deleteContainer(id: number): void {
    this.apiService.deleteContainer(id).subscribe({
      next: () => {
        // Убираем из локального массива
        this.containers = this.containers.filter(c => c.id !== id);
      },
      error: (err) => console.error('Failed to delete container', err)
    });
  }
}
