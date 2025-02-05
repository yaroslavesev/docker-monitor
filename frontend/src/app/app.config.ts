import { Routes } from '@angular/router';
import { ContainersComponent } from './pages/containers/containers.component';
import { provideHttpClient } from '@angular/common/http';

const appRoutes: Routes = [
  { path: '', component: ContainersComponent }
];

export const appConfig = {
  providers: [provideHttpClient()],
  routes: appRoutes
};
