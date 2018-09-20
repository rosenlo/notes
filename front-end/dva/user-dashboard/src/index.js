import dva from 'dva';
// import './index.css';
import '../index.html'
import createLoading from 'dva-loading';
import {services} from './utils/config'


// 1. Initialize
const app = dva();


// 2. Plugins
app.use(createLoading());

// 3. Model
function registerModel(modelList) {
  modelList.map(model => {
    app.model(require('./models/' + model))
  })
}
registerModel(['app']);
registerModel(services.map(service => service));

// app.model(require('./models/example'));

// 4. Router
app.router(require('./router'));

// 5. Start
app.start('#root');
