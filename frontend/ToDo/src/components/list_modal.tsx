import React from "react";
import { Button, Text, View } from "react-native";
import List from "../models/list";

interface Props {
    list: List;
}

const ListModal: React.FC<Props> = ({ list }) => {
    return (
        <View>
            <Text>
                {list.name}
            </Text>
        </View>
    );
}

export default ListModal;