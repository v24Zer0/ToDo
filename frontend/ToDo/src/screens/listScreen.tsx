import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";
import UserLists from "../components/user_lists";
import RootStackParamList from "./rootStackParamList";

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const ListScreen = ({ navigation, route }: Props) => {
    return (
        <UserLists />
    );
}

export default ListScreen;