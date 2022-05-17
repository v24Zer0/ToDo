import React from "react";
import { Button, View } from "react-native";
import List from "../models/list";

interface Props {
    list: List;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
}

const ListModal: React.FC<Props> = ({ list, setModalVisible }) => {
    return (
        <View>
            {list.name}
        </View>
    );
}

export default ListModal;