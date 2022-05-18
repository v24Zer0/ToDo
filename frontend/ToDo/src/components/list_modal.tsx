import React from "react";
import { Button, View } from "react-native";
import List from "../models/list";

interface Props {
    list: List;
}

const ListModal: React.FC<Props> = ({ list }) => {
    return (
        <View>
            {list.name}
        </View>
    );
}

export default ListModal;