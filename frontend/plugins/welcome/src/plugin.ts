import { createPlugin } from '@backstage/core';
import Login from './components/Login';
import CreateUser from './components/Users';
import ResultPage from './components/Result';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/login', Login);
    router.registerRoute('/user', CreateUser);
    router.registerRoute('/results', ResultPage);
  },
});
