/*
 * www
 * Copyright (C) 2025  flowerinsnow
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package cn.flowerinsnow.www.message.error.badrequest;

import cn.flowerinsnow.www.message.error.ErrorMessage;

public record RequestParameterMissing(String param) implements ErrorMessage {
    @Override
    public String string() {
        return "缺少请求参数：" + param;
    }

    public static RequestParameterMissing ofName(String param) {
        return new RequestParameterMissing(param);
    }
}
