import List from "../models/list";
import User from "../models/user";

type RootStackParamList = {
    Login: undefined;
    List: undefined;
    Item: {list: List}; 
    User: {user: User}
}

export default RootStackParamList;