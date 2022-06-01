import List from "../models/list";
import User from "../models/user";

type RootStackParamList = {
    Login: undefined;
    List: undefined;
    Item: {list: List}; 
    User: undefined;
    Signup: undefined;
}

export default RootStackParamList;