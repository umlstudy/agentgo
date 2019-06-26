import { createStore } from 'redux';
import modules from './modules';

const configure = () => {
    // const store = createStore(modules);

    // 개발을 더 편하게 하기 위해서 redux-devtools 라는 
    // 크롬 익스텐션을 사용해볼건데요, 이를 사용하기 위해선 크롬 웹스토어 에서 설치를 하고, 스토어 생성 함수를 조금 바꿔주어야 합니다.
    const devTools = (window as any).__REDUX_DEVTOOLS_EXTENSION__ && (window as any).__REDUX_DEVTOOLS_EXTENSION__()
    const store = createStore(modules, devTools);

    return store;
}

export default configure;