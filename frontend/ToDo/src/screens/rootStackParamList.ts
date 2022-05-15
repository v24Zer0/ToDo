import List from "../models/list";

type RootStackParamList = {
    Login: undefined;
    List: undefined;
    Item: {list: List}; 
}

export default RootStackParamList;