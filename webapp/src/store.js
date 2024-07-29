// store.js
import { createStore, applyMiddleware } from 'redux';
import rootReducer from './reducers'; // Import your root reducer

const store = createStore(rootReducer, applyMiddleware(/* middleware if any */));

export default store;
