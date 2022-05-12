import React from "react";
import { Text, View } from "react-native";
import List from "../models/list";

interface ListProps {
    list: List;
    // setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    // setModalList: React.Dispatch<React.SetStateAction<List>>;
}

// Add Pressable component to trigger modal display
const ListComponent: React.FC<ListProps> = ({ list }) => {
    return (
        <View>
            <Text>
                {list.name}
            </Text>
        </View>
    );
}

export default ListComponent